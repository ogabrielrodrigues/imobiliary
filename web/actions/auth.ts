'use server'

import { sign_schema } from "@/app/(auth)/cadastro/_components/sign-form"
import { env } from "@/lib/env"
import { User } from "@/types/user"
import jwt, { JwtPayload } from "jsonwebtoken"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import { z } from "zod"

export async function token(): Promise<string | undefined> {
  const cookiesStore = await cookies()

  const hasUser = cookiesStore.has("imobiliary-user")

  if (!hasUser) {
    return undefined
  }

  const token = cookiesStore.get("imobiliary-user")!.value

  return token
}

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
  const response = await fetch(`${env.SERVER_ADDR}/users/auth`, {
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

export async function sign(data: z.infer<typeof sign_schema>): Promise<number> {
  const response = await fetch(`${env.SERVER_ADDR}/users`, {
    method: "POST",
    body: JSON.stringify({
      fullname: data.fullname,
      creci_id: data.creci_id,
      email: data.email,
      cellphone: data.cellphone,
      password: data.password,
    }),
  })

  if (response.status !== 201) {
    return response.status
  }

  redirect("/login")
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete("imobiliary-user")

  redirect("/login")
}
