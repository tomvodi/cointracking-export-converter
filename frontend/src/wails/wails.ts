import {WailsApi} from "@/wails/wails_api";
import {GetExportFiles, OpenExportFile} from "@wails/go/cointracking/ct";
import {ExportToBlockpitXlsx} from "@wails/go/blockpit/bp";
import {common} from "@wails/go/models";
import {
    AllTimezones,
    BlockpitTxTypes,
    SetCointracking2BlockpitMapping,
    SetSwapHandling,
    SetTimezone,
    SwapHandling,
    Timezone,
    TxTypeMappings,
} from "@wails/go/config/appConfig";

export class Wails implements WailsApi {

    OpenExportFile(timezone: string): Promise<string> {
        return OpenExportFile(timezone);
    }

    ExportToBlockpitXlsx(): Promise<void> {
        return ExportToBlockpitXlsx();
    }

    AllTimezones(): Promise<Array<common.TimezoneData>> {
        return AllTimezones();
    }

    BlockpitTxTypes(): Promise<Array<common.TxDisplayName>> {
        return BlockpitTxTypes();
    }

    GetExportFiles(): Promise<Array<common.ExportFileInfo>> {
        return GetExportFiles();
    }

    SetCointracking2BlockpitMapping(ctType: string, bpType: string): Promise<void> {
        return SetCointracking2BlockpitMapping(ctType, bpType);
    }

    SetSwapHandling(swapHandling: string): Promise<void> {
        return SetSwapHandling(swapHandling);
    }

    SetTimezone(timezone: string): Promise<void> {
        return SetTimezone(timezone);
    }

    SwapHandling(): Promise<string> {
        return SwapHandling();
    }

    Timezone(): Promise<string> {
        return Timezone();
    }

    TxTypeMappings(): Promise<Array<common.Ct2BpTxMapping>> {
        return TxTypeMappings();
    }
}