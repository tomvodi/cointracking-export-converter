import {WailsRuntimeApi} from "@/wails/wails_runtime_api";
import {EventsOn} from "@wails/runtime";

export class WailsRuntime implements WailsRuntimeApi {
    EventsOn(eventName: string, callback: (...data: any) => void): () => void {
        return EventsOn(eventName, callback);
    }
}