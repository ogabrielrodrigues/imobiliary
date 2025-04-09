import { users } from "@/mock/users";
import { cookies } from "next/headers";

import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
  const body = await request.json()

  const user = users.find(user => user.email === body.email && user.password === body.password)

  if (!user) {
    return new NextResponse(null, { status: 400 })
  }

  const cookieStore = await cookies()
  const imobiliaryUser = cookieStore.set("imobiliary-user", JSON.stringify(user), {
    maxAge: 60 * 60 * 24 * 30,
    path: "/",
  })

  return new NextResponse(null, { 
      status: 200,
      headers: {
        "Set-Cookie": `imobiliary-user=${imobiliaryUser}`
      } 
    }
  )
}