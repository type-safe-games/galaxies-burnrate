export namespace data {
	
	export class SaveMeta {
	    save_id: string;
	    captain_name: string;
	    ship_name: string;
	    credits: number;
	    current_day: number;
	
	    static createFrom(source: any = {}) {
	        return new SaveMeta(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.save_id = source["save_id"];
	        this.captain_name = source["captain_name"];
	        this.ship_name = source["ship_name"];
	        this.credits = source["credits"];
	        this.current_day = source["current_day"];
	    }
	}

}

