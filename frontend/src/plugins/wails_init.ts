import {App, Plugin} from "vue";
import {Wails} from "@/wails/wails";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey, wailsRuntimeInjKey} from "@/injection_keys";
import {WailsRuntime} from "@/wails/wails_runtime";
import {common} from "@wails/go/models";
import {useApplicationStore} from "@/stores/application_store";

const setExportFiles = async (files: Array<common.ExportFileInfo>) => {
    const appStore = useApplicationStore();
    appStore.exportFiles = files
}

/*
    * This plugin initializes the Wails API client and runtime.
    * It provides the Wails API client and runtime as injection keys.
    * It also listens for the ExportFilesChanged event and updates the exportFiles store.
 */
export const WailsInit: Plugin = {
    install(app: App) {
        const wailsApiClient = new Wails()
        app.provide<WailsApi>(wailsClientInjKey, wailsApiClient)
        const wailsRuntime = new WailsRuntime()
        app.provide<WailsRuntime>(wailsRuntimeInjKey, wailsRuntime)

        wailsRuntime.EventsOn("ExportFilesChanged", setExportFiles)
    }
}