'use server'

import { env } from "@/lib/env"
import { token } from "./auth"

export async function updateAvatar(formData: FormData): Promise<{ status: number, url: string }> {
  const auth_token = await token()

  if (!auth_token) {
    return { status: 401, url: "" }
  }

  const response = await fetch(`${env.SERVER_ADDR}/users/avatar`, {
    method: "POST",
    body: formData,
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  const data = await response.json()

  return { status: response.status, url: data.url }
}