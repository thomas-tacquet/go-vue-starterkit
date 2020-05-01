const path = require('path');
const webpack = require('webpack');
const VueLoader = require('vue-loader');
const VuetifyLoaderPlugin = require('vuetify-loader/lib/plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer')
    .BundleAnalyzerPlugin;
const HardSourceWebpackPlugin = require('hard-source-webpack-plugin');

const productionURL = 'http://localhost:8081';
const developmentURL = 'http://localhost:8080';

const URL =
    process.env.NODE_ENV === 'production' ? productionURL : developmentURL;

module.exports = {
    entry: './frontend/main.js',
    mode: 'development',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: '[name].[contenthash:8].js',
        publicPath: '/assets/'
        //sourceMapFilename: 'dist/build.map'
    },
    // To enable using @XXX in other frontend files
    resolve: {
        extensions: ['.js', '.vue', '.json', '.ts'],
        alias: {
            '@Util': path.resolve(__dirname, 'frontend', 'utils'),
            '@Compo': path.resolve(__dirname, 'frontend', 'components'),
            '@Views': path.resolve(__dirname, 'frontend', 'views'),
            '@Layout': path.resolve(__dirname, 'frontend', 'layouts'),
            '@Asset': path.resolve(__dirname, 'frontend', 'assets')
        }
    },
    plugins: [
        new VuetifyLoaderPlugin(),
        new VueLoader.VueLoaderPlugin(),
        new HtmlWebpackPlugin({
            template: './frontend/index.html'
        }),
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery'
        }),
        new BundleAnalyzerPlugin({
            analyzerMode: 'static',
            reportFilename: '../out/bundle_report.html',
            openAnalyzer: false
        }),
        new HardSourceWebpackPlugin({
            cachePrune: {
                maxAge: 12 * 60 * 60 * 1000,
                sizeThreshold: 75 * 1024 * 1024
            }
        })
    ],
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: {
                    loader: 'vue-loader'
                }
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        cacheDirectory: true
                    }
                }
            },
            {
                test: /\.ts$/,
                exclude: /node_modules/,
                loader: 'ts-loader'
            },
            {
                test: /\.(png|jpg|gif|svg)$/,
                use: {
                    loader: 'file-loader',
                    query: {
                        name: './img/[name]-[hash:8].[ext]'
                    }
                }
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.(ttf|TTF|eot|otf|woff|woff2)$/,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: './fonts/[name].[ext]'
                    }
                }
            },
            {
                test: /\.s(c|a)ss$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    {
                        loader: 'sass-loader',
                        options: {
                            implementation: require('sass'),
                            fiber: require('fibers'),
                            indentSyntax: true
                        }
                    }
                ]
            }
        ]
    },
    optimization: {
        runtimeChunk: 'single',
        splitChunks: {
            chunks: 'all',
            maxInitialRequests: Infinity,
            maxAsyncRequests: Infinity,
            minSize: 10000,
            cacheGroups: {
                vendor: {
                    test: /[\\/]node_modules[\\/]/,
                    name(module) {
                        // get the name. E.g. node_modules/packageName/not/this/part.js
                        // or node_modules/packageName
                        const packageName = module.context.match(
                            /[\\/]node_modules[\\/](.*?)([\\/]|$)/
                        )[1];

                        // npm package names are URL-safe, but some servers don't like @ symbols
                        return `npm.${packageName.replace('@', '')}`;
                    }
                }
            }
        }
    }
};

if (process.env.NODE_ENV === 'production') {
    module.exports.mode = 'production';
} else {
    module.exports.mode = 'development';
    module.exports.devtool = '#eval-source-map';
}

module.exports.plugins = (module.exports.plugins || []).concat([
    new webpack.DefinePlugin({
        API_URL: JSON.stringify(URL)
    })
]);
