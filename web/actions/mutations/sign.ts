'use server'

import { sign_schema } from "@/app/(auth)/cadastro/_components/sign-form"
import { env } from "@/lib/env"
import { z } from "zod"

export async function sign(values: z.infer<typeof sign_schema>): Promise<number> {
  try {
    const response = await fetch(`${env.SERVER_ADDR}/users`, {
      method: "POST",
      body: JSON.stringify({
        fullname: values.fullname,
        creci_id: values.creci_id,
        email: values.email,
        cellphone: values.cellphone,
        password: values.password,
      }),
    })

    return response.status
  } catch {
    return 500
  }
}