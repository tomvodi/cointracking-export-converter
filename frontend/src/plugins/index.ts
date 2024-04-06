import {App} from "vue";
import {router} from "@/router";
import vuetify from "@/plugins/vuetify";
import {PiniaPlugin} from "@/plugins/pinia";
import {WailsInit} from "@/plugins/wails_init";

export function registerPlugins(app: App) {
    app.use(router)
    app.use(vuetify)
    app.use(PiniaPlugin)
    app.use(WailsInit)
}