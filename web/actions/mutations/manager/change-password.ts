"use server"

import type { ResetPasswordRequest } from "@/components/reset-password-dialog"
import { env } from "@/lib/env"

export async function changePassword(
  data: ResetPasswordRequest,
): Promise<number> {
  try {
    const response = await fetch(`${env.SERVER_ADDR}/manager/reset`, {
      method: "PUT",
      body: JSON.stringify(data),
    })

    return response.status
  } catch {
    return 500
  }
}
