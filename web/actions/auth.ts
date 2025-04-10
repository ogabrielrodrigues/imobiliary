'use server'

import { env } from "@/lib/env"
import { User } from "@/types/user"
import jwt, { JwtPayload } from "jsonwebtoken"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

export async function auth(): Promise<User | undefined> {
  const cookiesStore = await cookies()

  const hasUser = cookiesStore.has("imobiliary-user")

  if (!hasUser) {
    return undefined
  }

  const token = jwt.verify(cookiesStore.get("imobiliary-user")!.value, env.JWT_SECRET!) as JwtPayload

  return token.user
}

export async function login(email: string, password: string): Promise<number> {
  const response = await fetch(`${env.SERVER_ADDR}/auth`, {
    method: "POST",
    body: JSON.stringify({ email, password }),
  })

  if (response.status !== 200) {
    return response.status
  }

  const token = response.headers.get("Authorization")!.split(" ")[1]

  const cookieStore = await cookies()
  cookieStore.set("imobiliary-user", token, {
    maxAge: 60 * 60 * 24 * 30,
    path: "/",
  })

  redirect("/dashboard")
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete("imobiliary-user")

  redirect("/login")
}
