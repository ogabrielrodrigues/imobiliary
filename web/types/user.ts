import { z } from "zod";

export const user_schema = z.object({
  id: z.string().uuid(),
  fullname: z.string(),
  creci_id: z.string().min(7).max(7),
  email: z.string().email(),
  password: z.string().min(8),
  avatar: z.string().url().optional(),
  cellphone: z.string().min(12).max(13),
  plan: z.object({
    id: z.string().uuid(),
    kind: z.enum(["free", "pro"]),
    price: z.number(),
    properties_total_quota: z.number(),
    properties_used_quota: z.number(),
    properties_remaining_quota: z.number(),
  })
});

export type User = z.infer<typeof user_schema>;

