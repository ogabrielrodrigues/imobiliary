'use server'

import { SignRequest } from "@/app/(auth)/cadastro/_components/sign-form"
import { env } from "@/lib/env"

export async function sign(values: SignRequest): Promise<number> {
  try {
    const response = await fetch(`${env.SERVER_ADDR}/manager`, {
      method: "POST",
      body: JSON.stringify({
        fullname: values.fullname,
        phone: values.phone,
        email: values.email,
        password: values.password,
      }),
    })

    return response.status
  } catch {
    return 500
  }
}