'use server'

import { LoginRequest } from "@/app/(auth)/login/_components/login-form"
import { env } from "@/lib/env"
import { cookies } from "next/headers"

export async function login(values: LoginRequest): Promise<number> {
  try {
    const response = await fetch(`${env.SERVER_ADDR}/auth`, {
      method: "POST",
      body: JSON.stringify(values),
    })

    if (response.status !== 200) {
      return response.status
    }

    const token = response.headers.get("Authorization")!.split(" ")[1]

    const cookieStore = await cookies()
    cookieStore.set("imobiliary-user", token, {
      maxAge: 60 * 60 * 24 * 30,
      path: "/",
      httpOnly: true,
    })

    return response.status
  } catch {
    return 500
  }
}