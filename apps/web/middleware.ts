import { cookies } from "next/headers"
import { NextRequest, NextResponse } from "next/server"

const protectedRoutes = '/dashboard'
const publicRoutes = ['/login', '/cadastro']

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname
  const isProtectedRoute = path.includes(protectedRoutes)
  const isPublicRoute = publicRoutes.includes(path)

  const cookieStore = await cookies()
  const hasUser = cookieStore.has("imobiliary-user")

  if (isProtectedRoute && !hasUser) {
    return NextResponse.redirect(new URL('/login', req.nextUrl))
  }

  if (isPublicRoute && hasUser) {
    return NextResponse.redirect(new URL('/dashboard', req.nextUrl))
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}