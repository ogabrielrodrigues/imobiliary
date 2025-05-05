'use server'

import { env } from "@/lib/env"
import { User } from "@/types/user"

type GetUserResponse = { user: User | undefined, status: number }

export async function getUser(id: string): Promise<GetUserResponse> {
  try {
    const response = await fetch(`${env.SERVER_ADDR}/users/${id}`, {
      method: "GET",
    })

    if (response.status !== 200) {
      return { user: undefined, status: response.status }
    }

    const user = await response.json() as User

    return { user, status: response.status }
  } catch {
    return { user: undefined, status: 500 }
  }
}