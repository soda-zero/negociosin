import type { ZodSchema } from "zod";

export function validateForm(schema: ZodSchema, form: {}) {
  const result = schema.safeParse(form);
  if (result.success === false) {
    return {
      errors: result.error.formErrors.fieldErrors,
      valid: false,
    };
  }
  return { errors: {}, valid: true };
}
