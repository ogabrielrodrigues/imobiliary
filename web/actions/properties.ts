'use server'

import { env } from "@/lib/env"
import { Property } from "@/types/property"
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

export async function getProperty(id: string): Promise<Property[]> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/properties/${id}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  return response.json()
}