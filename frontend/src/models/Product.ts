import { z } from "zod";

const errors = {
  obligatory: "Campo obligatorio.",
  name: "Debe tener al menos 2 caractéres.",
  non_negative: "No puede ser un número negativo.",
  quantity: "No puede ser menor a 1.",
  num: "Deber ser un número.",
};
export const ProductSchema = z.object({
  name: z
    .string({
      required_error: errors.obligatory,
    })
    .min(2, errors.name),
  iva: z.number(),
  provider_id: z.number().nullable(),
  original_cost_price: z
    .number({
      required_error: errors.obligatory,
      invalid_type_error: errors.num,
    })
    .nonnegative({ message: errors.non_negative }),
  quantity: z
    .number({ required_error: errors.obligatory })
    .min(1, errors.quantity),
  internal_tax: z.number().min(0, errors.non_negative).optional(),
  category_id: z.number().nullable(),
});

export const ProductKeys = ProductSchema.keyof().Enum;
export type ProductKey = keyof typeof ProductKeys;
export type Product = z.infer<typeof ProductSchema>;
