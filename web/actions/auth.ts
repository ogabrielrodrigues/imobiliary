'use server'

import { User } from "@/types/user"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

export async function auth(): Promise<User | undefined> {
  const cookiesStore = await cookies()

  const hasUser = cookiesStore.has("imobiliary-user")

  if (!hasUser) {
    return undefined
  }

  return JSON.parse(cookiesStore.get("imobiliary-user")!.value)
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete("imobiliary-user")

  redirect("/login")
}
