import {createPinia, PiniaVuePlugin} from "pinia";
import {App} from "vue";


export const PiniaPlugin = {
    install(app: App) {
        app.use(createPinia())
        app.use(PiniaVuePlugin)
    }
}