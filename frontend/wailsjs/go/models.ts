export namespace backend {
	
	export class Category {
	    id?: number;
	    name?: string;
	    profit_percent?: number;
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.profit_percent = source["profit_percent"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
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
	export class Product {
	    id?: number;
	    name?: string;
	    original_cost_price?: number;
	    unit_cost_price?: number;
	    unit_sell_price?: number;
	    category_id?: number;
	    provider_id?: number;
	    quantity?: number;
	    iva?: number;
	    internal_tax?: number;
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Product(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.original_cost_price = source["original_cost_price"];
	        this.unit_cost_price = source["unit_cost_price"];
	        this.unit_sell_price = source["unit_sell_price"];
	        this.category_id = source["category_id"];
	        this.provider_id = source["provider_id"];
	        this.quantity = source["quantity"];
	        this.iva = source["iva"];
	        this.internal_tax = source["internal_tax"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
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
	export class ProductWithCategoryAndProvider {
	    id?: number;
	    name?: string;
	    original_cost_price?: number;
	    unit_cost_price?: number;
	    unit_sell_price?: number;
	    category_id?: number;
	    provider_id?: number;
	    quantity?: number;
	    iva?: number;
	    internal_tax?: number;
	    category_profit_percent?: number;
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	    category_name?: string;
	    provider?: string;
	
	    static createFrom(source: any = {}) {
	        return new ProductWithCategoryAndProvider(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.original_cost_price = source["original_cost_price"];
	        this.unit_cost_price = source["unit_cost_price"];
	        this.unit_sell_price = source["unit_sell_price"];
	        this.category_id = source["category_id"];
	        this.provider_id = source["provider_id"];
	        this.quantity = source["quantity"];
	        this.iva = source["iva"];
	        this.internal_tax = source["internal_tax"];
	        this.category_profit_percent = source["category_profit_percent"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.category_name = source["category_name"];
	        this.provider = source["provider"];
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
	export class Provider {
	    id?: number;
	    name?: string;
	    phone_number?: string;
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Provider(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.phone_number = source["phone_number"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
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

