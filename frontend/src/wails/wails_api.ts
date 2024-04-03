import {common} from "@wails/go/models";

export interface WailsApi {
    // CoinTracking methods
    OpenExportFile(timezone: string): Promise<string>;

    GetExportFiles(): Promise<Array<common.ExportFileInfo>>;

    // Blockpit methods
    ExportToBlockpitXlsx(): Promise<void>;

    // App Config methods
    AllTimezones(): Promise<Array<common.TimezoneData>>;

    BlockpitTxTypes(): Promise<Array<common.TxDisplayName>>;

    SetCointracking2BlockpitMapping(ctType: string, bpType: string): Promise<void>;

    SetSwapHandling(swapHandling: string): Promise<void>;

    SetTimezone(timezone: string): Promise<void>;

    SwapHandling(): Promise<string>;

    Timezone(): Promise<string>;

    TxTypeMappings(): Promise<Array<common.Ct2BpTxMapping>>;
}