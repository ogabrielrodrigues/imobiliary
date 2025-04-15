'use server'

import { env } from "@/lib/env"
import { token } from "./auth"

export async function updateAvatar(formData: FormData): Promise<number> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/users/avatar`, {
    method: "POST",
    body: formData,
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  return response.status
}