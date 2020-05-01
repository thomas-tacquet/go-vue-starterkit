<template>
    <div id="app">
        <v-app>
            <component :is="layout">
                <router-view/>
            </component>
        </v-app>
    </div>
</template>

<script>
    import {GetLayout} from '@Layout'

    export default {
        name: 'app',
        computed: {
            layout() {
                return GetLayout(this.$route.meta.layout)
            }
        },
    };
</script>

