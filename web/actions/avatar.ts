'use server'

import { env } from "@/lib/env"
import { token } from "./auth"
import { cookies } from "next/headers"

export async function updateAvatar(formData: FormData): Promise<number> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/users/avatar`, {
    method: "POST",
    body: formData,
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  const new_token = response.headers.get("Authorization")!.split(" ")[1]

  const cookieStore = await cookies()
  cookieStore.set("imobiliary-user", new_token, {
    maxAge: 60 * 60 * 24 * 30,
    path: "/",
  })

  return response.status
}