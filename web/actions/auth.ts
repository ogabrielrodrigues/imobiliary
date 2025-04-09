'use server'

import { User } from "@/types/user"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

export async function auth(): Promise<User | undefined> {
  const cookiesStore = await cookies()

  const user = cookiesStore.get("imobiliary-user")
  if (!user) {
    return undefined
  }

  return JSON.parse(user.value)
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete("imobiliary-user")

  redirect("/login")
}
