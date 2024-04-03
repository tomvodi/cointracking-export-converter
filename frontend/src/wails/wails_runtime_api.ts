export interface WailsRuntimeApi {
    EventsOn(eventName: string, callback: (...data: any) => void): () => void;
}