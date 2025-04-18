import { z } from "zod";

const plan_schema = z.object({
  id: z.string().uuid(),
  kind: z.enum(["free", "pro"]),
  price: z.number(),
  properties_total_quota: z.number(),
  properties_used_quota: z.number(),
  properties_remaining_quota: z.number(),
})

export type Plan = z.infer<typeof plan_schema>