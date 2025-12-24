import { type NextRequest, NextResponse } from "next/server"
import { auth } from "./actions/queries/auth/auth"

const protectedRoutes = ["/dashboard"]
const publicRoutes = ["/login", "/cadastro"]

export default async function proxy(req: NextRequest) {
	const path = req.nextUrl.pathname
	const isProtectedRoute = protectedRoutes.some(route => path.startsWith(route))
	const isPublicRoute = publicRoutes.some(route => path.startsWith(route))

	const isAutheticated = await auth()

	if (isProtectedRoute && !isAutheticated) {
		return NextResponse.redirect(new URL("/login", req.nextUrl))
	}

	if (isPublicRoute && isAutheticated) {
		return NextResponse.redirect(new URL("/dashboard", req.nextUrl))
	}

	return NextResponse.next()
}

export const config = {
	matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)"],
}
