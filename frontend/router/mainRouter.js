import Vue from 'vue'
import VueRouter from 'vue-router'

import {Layouts} from '@Layout'

Vue.use(VueRouter);

export default new VueRouter({
    mode: 'history',
    base: __dirname,
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('@Views/Home'),
            meta: {layout: Layouts.DEFAULT}
        },
        {
            path: '*',
            redirect: '/'
        }
    ]
});
