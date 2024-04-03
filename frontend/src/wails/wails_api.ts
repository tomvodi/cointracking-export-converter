export interface WailsApi {
    OpenExportFile(timezone: string): Promise<string>;
}