import { z } from "zod";

const errors = {
  obligatory: "Campo obligatorio.",
  name: "Debe tener al menos 2 caractéres.",
  non_negative: "No puede ser un número negativo.",
  num: "Deber ser un número.",
};
export const CategorySchema = z.object({
  name: z
    .string({
      required_error: errors.obligatory,
    })
    .min(2, errors.name),
  profit_percent: z
    .number({
      required_error: errors.obligatory,
      invalid_type_error: errors.num,
    })
    .nonnegative({ message: errors.non_negative }),
});

export const CategoryKeys = CategorySchema.keyof().Enum;
export type CategoryKey = keyof typeof CategoryKeys;
export type Category = z.infer<typeof CategorySchema>;
