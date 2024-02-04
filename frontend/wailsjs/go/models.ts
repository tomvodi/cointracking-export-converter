export namespace common {
	
	export class TxDisplayName {
	    title: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new TxDisplayName(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.value = source["value"];
	    }
	}
	export class Ct2BpTxMapping {
	    cointracking: TxDisplayName;
	    blockpit: TxDisplayName;
	
	    static createFrom(source: any = {}) {
	        return new Ct2BpTxMapping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cointracking = this.convertValues(source["cointracking"], TxDisplayName);
	        this.blockpit = this.convertValues(source["blockpit"], TxDisplayName);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

