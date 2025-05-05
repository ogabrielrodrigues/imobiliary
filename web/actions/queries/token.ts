import { cookies } from "next/headers"

export async function token(): Promise<string | undefined> {
  const cookiesStore = await cookies()

  if (!cookiesStore.has("imobiliary-user")) {
    return undefined
  }

  const token = cookiesStore.get("imobiliary-user")!.value

  return token
}