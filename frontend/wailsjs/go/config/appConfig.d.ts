// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {common} from '../models';

export function AllTimezones():Promise<Array<common.TimezoneData>>;

export function BlockpitTxTypes():Promise<Array<common.TxDisplayName>>;

export function SetCointracking2BlockpitMapping(arg1:string,arg2:string):Promise<void>;

export function SetSwapHandling(arg1:string):Promise<void>;

export function SetTimezone(arg1:string):Promise<void>;

export function SwapHandling():Promise<string>;

export function Timezone():Promise<string>;

export function TxTypeMappings():Promise<Array<common.Ct2BpTxMapping>>;
