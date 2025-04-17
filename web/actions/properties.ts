'use server'

import { property_schema } from "@/app/dashboard/locacao/imoveis/_components/new-property-form"
import { searchCEP } from "@/lib/cep/api-brasil"
import { env } from "@/lib/env"
import { Property } from "@/types/property"
import { z } from "zod"
import { token } from "./auth"

export async function getProperties(): Promise<Property[]> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/properties`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  return response.json()
}

export async function getProperty(id: string): Promise<Property> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/properties/${id}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  return response.json()
}

export async function createProperty(data: z.infer<typeof property_schema>): Promise<number> {
  const auth_token = await token()

  const cep = data.address.zip_code.replace("-", "")
  const found = await searchCEP(env.CEP_API_ADDR!, cep)

  if (found == undefined) {
    return 400
  }

  data.address.state = found.state

  const response = await fetch(`${env.SERVER_ADDR}/properties`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  return 200
}