import { z } from "zod";

export const user_schema = z.object({
  id: z.string().uuid(),
  fullname: z.string(),
  creci_id: z.string().min(7).max(7),
  email: z.string().email(),
  password: z.string().min(8),
  avatar: z.string().url(),
  cellphone: z.string().min(12).max(13)
});

export type User = z.infer<typeof user_schema>;

