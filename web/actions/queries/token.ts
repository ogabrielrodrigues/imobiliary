import { cookies } from "next/headers"

export async function token(): Promise<string> {
  const cookiesStore = await cookies()

  const token = cookiesStore.get("imobiliary-user")!.value

  return token
}