import { cookies } from 'next/headers'
import { NextRequest, NextResponse } from 'next/server'

const protectedRoutes = ['/dashboard']
const rootRoute = '/'
const publicRoutes = ['/login', '/cadastro']

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname
  const isProtectedRoute = protectedRoutes.includes(path)
  const isPublicRoute = publicRoutes.includes(path)

  const cookieStore = await cookies()
  const imobiliaryUser = cookieStore.get("imobiliary-user")
  const user = imobiliaryUser ? JSON.parse(imobiliaryUser.value) : null

  if (isProtectedRoute && !user) {
    return NextResponse.redirect(new URL('/login', req.nextUrl))
  }

  if (isPublicRoute && user) {
    return NextResponse.redirect(new URL('/dashboard', req.nextUrl))
  }

  // if (
  //   isPublicRoute &&
  //   user &&
  //   !req.nextUrl.pathname.startsWith('/dashboard')
  // ) {
  //   return NextResponse.redirect(new URL('/dashboard', req.nextUrl))
  // }

  const headers = new Headers()
  headers.set('x-pathname', path)


  return NextResponse.next({ headers })
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}