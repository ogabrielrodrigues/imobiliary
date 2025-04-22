import { z } from "zod";

export const property_schema = z.object({
  id: z.string().uuid(),
  water_id: z.string(),
  energy_id: z.string(),
  user_id: z.string().uuid(),
  owner_id: z.string().uuid(),
  kind: z.enum(['Residencial', 'Comercial', 'Industrial', 'Terreno', 'Rural']),
  status: z.enum(['Disponível', 'Ocupado', 'Indisponível', 'Reservado', 'Reformando']),
  address: z.object({
    full_address: z.string(),
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

export type Property = z.infer<typeof property_schema>;

