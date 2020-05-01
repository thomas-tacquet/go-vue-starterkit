const gulp = require('gulp');
const del = require('del');
const webpack = require('webpack');
const webpackStream = require('webpack-stream');
const webpackConfig = require('./webpack.config.js');
const livereload = require('gulp-livereload');
const fs = require('fs');
const path = require('path');

const util = require('gulp-util'),
    notifier = require('node-notifier'),
    child = require('child_process'),
    os = require('os');

let server = null;

gulp.task('clean', function clean() {
    return del(['./out/**/*', './dist/**/*', './backend/debug'])
});

gulp.task('build:front', function buildFront(done) {
    return gulp.src('./frontend/main.js')
        .pipe(webpackStream(webpackConfig, webpack))
        .on('error', function handleError() {
            this.emit('end'); // Recover from errors
        })
        .on('end', function endPipe() {
            done();
        })
        .pipe(gulp.dest('./dist/'))
        .pipe(livereload({
            key: fs.readFileSync(path.join(__dirname, './tls/server.key'), 'utf-8'),
            cert: fs.readFileSync(path.join(__dirname, './tls/server.pem'), 'utf-8')
        }));

});

gulp.task('build:back', function buildBack(done) {
    child.spawnSync('make', [], {cwd: './backend/'});

    // Build application
    const env = Object.create(process.env);
    const build = child.spawnSync('go', ['build', '-o', '../out/backend'], {cwd: './backend/', env: env});

    if (build.stderr && build.stderr.length) {
        util.log(util.colors.red('Something wrong :'));
        const lines = build.stderr.toString()
            .split('\n').filter(function (line) {
                return line.length
            });
        for (let l in lines)
            util.log(util.colors.red(
                'Error (go build): ' + lines[l]
            ));
        notifier.notify({
            title: 'Error (go build)',
            message: lines
        });
    }
    done();

    return build;

});

gulp.task('watch:front', function watchFront() {
    livereload.listen();
    return gulp.watch('./frontend/**/*', gulp.series('build:front'));
});

gulp.task('watch:back', function () {
    livereload.listen();
    return gulp.watch([
        'backend/*.go',
        'backend/**/*.go',
        '!backend/tests/**/*.go',
        '!backend/**/*_test.go',
        'conf/**/*'
    ], gulp.series(
        'build:back',
        'spawn:gin'
    ));
});

gulp.task('watch:test', function watchTest() {
    child.spawnSync('npm', ['run', 'test'], {stdio: 'inherit'});

    return gulp.watch(['backend/*.go', 'backend/**/*.go'], function () {
        return child.spawnSync('npm', ['run', 'test'], {stdio: 'inherit'})
    })
});

// start backend
gulp.task('spawn:gin', function spawnGin(done) {

    // Stop the server
    if (server && server !== 'null') {
        server.kill();
    }

    // Run the server
    if (os.platform() === 'win32') {
        server = child.spawn('out/backend.exe')
            .on('error', function (err) {
                process.stderr.write(err.toString())
            })
            .on('end', function () {
                done()
            });
    } else {
        server = child.spawn('out/backend')
            .on('error', function (err) {
                process.stderr.write(err.toString())
            })
            .on('end', function () {
                done()
            });
    }

    // Display terminal informations
    server.stderr.on('data', function (data) {
        process.stdout.write(data.toString());
    });
    server.stdout.on('data', function (data) {
        process.stdout.write(data.toString());
    });

    done();
});


function onSigTerm(signal) {
    if (server != null) {
        server.kill();
        server.on('exit', function () {
            child.spawnSync('kill', ['-s', 'SIGKILL', process.pid]);
        })
    } else {
        child.spawnSync('kill', ['-s', 'SIGKILL', process.pid]);
    }
}

process.on('SIGINT', onSigTerm); // CTRL+C
process.on('SIGTERM', onSigTerm);
process.on('SIGABRT', onSigTerm);
// SIGKILL non catcheable by gulp(see doc : https://nodejs.org/api/process.html)

// Main tasks
gulp.task('build', gulp.parallel('build:back', 'build:front'));
gulp.task('spawn', gulp.parallel('spawn:gin'));
gulp.task('watch', gulp.parallel('watch:back', 'watch:front'));
gulp.task('default', gulp.series('build', 'spawn', 'watch'));
gulp.task('debug', gulp.series('build:front', 'watch:front'));
