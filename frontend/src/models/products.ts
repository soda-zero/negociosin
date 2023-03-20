export type Product = {
  id?: number;
  name: string;
  cost_price: number;
  provider: string;
  created_at?: string;
  updated_at?: string;
};

export type Provider = {
  name?: string;
};
