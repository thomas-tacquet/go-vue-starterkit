import Vue from 'vue'
import vuetify from './plugins/vuetify'
import App from './App.vue'
import router from './router/mainRouter'
import VueTimeago from 'vue-timeago'
import vmodal from 'vue-js-modal'
import VueTheMask from 'vue-the-mask'

import 'core-js/modules/es.promise'
import 'core-js/modules/es.array.iterator'

require('typeface-roboto');

// Layouts
const DefaultLayout = () => import(/* webpackChunkName: 'layout' */ '@Layout/Default.vue');


Vue.component('default-layout', DefaultLayout);

Vue.use(vmodal);
Vue.use(VueTheMask);
Vue.use(VueTimeago, {
    name: 'Timeago',
    locale: 'fr',
    locales: {
        fr: require('date-fns/locale/fr')
    }
});

router.beforeEach((to, fm, next) => {
    next()
});

new Vue({
    vuetify,
    router,
    render: h => h(App)
}).$mount('#app');
