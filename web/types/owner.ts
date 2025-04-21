import { z } from "zod";

export const owner_schema = z.object({
  id: z.string().uuid(),
  fullname: z.string(),
  cpf: z.string().length(14),
  rg: z.string().max(15),
  email: z.string().email(),
  cellphone: z.string(),
  occupation: z.string(),
  marital_status: z.enum(["Solteiro(a)", "Casado(a)", "Amasiado(a)", "Divorciado(a)", "Viúvo(a)", "União Estável"]),
  manager_id: z.string().uuid(),
  address: z.object({
    full_address: z.string().optional(),
    mini_address: z.string(),
    street: z.string(),
    number: z.string(),
    neighborhood: z.string(),
    complement: z.string().optional(),
    city: z.string(),
    state: z.string(),
    state_abbr: z.string(),
    zip_code: z.string(),
  }),
});

export type Owner = z.infer<typeof owner_schema>;

