import { z } from "zod";

const errors = {
  obligatory: "Campo obligatorio.",
  name: "Debe tener al menos 2 caractéres.",
  phone: "El número de telefono debe tener al menos 8 caractéres.",
};
export const ProviderSchema = z.object({
  name: z
    .string({
      required_error: errors.obligatory,
    })
    .min(2, errors.name),
  phone_number: z
    .string({ invalid_type_error: "Tiene que ser un texto" })
    .min(8, { message: errors.phone })
    .nullish(),
});

export const ProviderKeys = ProviderSchema.keyof().Enum;
export type ProviderKey = keyof typeof ProviderKeys;
export type Provider = z.infer<typeof ProviderSchema>;
