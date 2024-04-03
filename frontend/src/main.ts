import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import '@mdi/font/css/materialdesignicons.css'

import 'vuetify/styles'
import {createVuetify} from 'vuetify'
import {aliases, mdi} from 'vuetify/iconsets/mdi'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {createPinia, PiniaVuePlugin} from "pinia";
import {router} from "./router";
import {Wails} from "@/wails/wails";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey, wailsRuntimeInjKey} from "@/injection_keys";
import {WailsRuntime} from "@/wails/wails_runtime";

const pinia = createPinia()

const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'light'
    },
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            mdi,
        },
    },
})

const app = createApp(App)
const wailsApiClient = new Wails()
app.provide<WailsApi>(wailsClientInjKey, wailsApiClient)
const wailsRuntime = new WailsRuntime()
app.provide<WailsRuntime>(wailsRuntimeInjKey, wailsRuntime)
app.use(router)
app.use(vuetify)
app.use(pinia)
app.use(PiniaVuePlugin)
app.mount('#app')
