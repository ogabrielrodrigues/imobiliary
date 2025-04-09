import { z } from "zod";

export const user_schema = z.object({
  id: z.string().uuid(),
  fullname: z.string(),
  creci: z.string().min(7).max(7),
  email: z.string().email(),
  password: z.string().min(8),
  avatar: z.string().url().optional(),
  telefone: z.string().min(12).max(13),
  plan: z.object({
    kind: z.enum(["free", "pro"]),
    propertiesTotalQuota: z.number(),
    propertiesUsedQuota: z.number(),
    propertiesRemainingQuota: z.number(),
  })
});

export type User = z.infer<typeof user_schema>;

