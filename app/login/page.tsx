import { Metadata } from "next"
import Link from "next/link"

import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"
import { UserAuthForm } from "@/components/user-auth-form"
import { Hexagon } from "lucide-react"
import { ModeToggle } from "@/components/mode-toggle"

export const metadata: Metadata = {
    title: "Autentificación - Metropolis",
    description:
        "La principal plataforma de empleo en el país. Nuestro portal de ofertas laborales y oportunidades de empleo cuenta con una amplia y completa cobertura en todo el territorio boliviano.",
}

export default function AuthenticationPage() {
    return (
        <>
            <div className="container relative h-screen flex-col items-center justify-center grid lg:max-w-none lg:grid-cols-2 lg:px-0">
                <div className="absolute right-4 top-4 md:right-8 md:top-8 flex items-center space-x-4">
                    <Link
                        href="/examples/authentication"
                        className={cn(buttonVariants({ variant: "ghost" }))}
                    >
                        Iniciar Sesión
                    </Link>
                    <ModeToggle variant="ghost" />
                </div>
                <div className="relative hidden h-full flex-col bg-muted p-10 text-white dark:border-r lg:flex">
                    <div className="absolute inset-0 bg-gradient-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90%" />
                    <div className="relative z-20 flex items-center text-lg font-medium">
                        <Hexagon className="mr-2 h-6 w-6" />
                        Metropolis
                    </div>
                    <div className="relative z-20 mt-auto">
                        <blockquote className="space-y-2">
                            <p className="text-lg">
                                &ldquo;La principal plataforma de empleo en el
                                país. Nuestro portal de ofertas laborales y
                                oportunidades de empleo cuenta con una amplia y
                                completa cobertura en todo el territorio
                                boliviano.&rdquo;
                            </p>
                            <footer className="text-sm">— Sofia Davis</footer>
                        </blockquote>
                    </div>
                </div>
                <div className="lg:p-8">
                    <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
                        <div className="flex flex-col space-y-2 text-center">
                            <h1 className="text-2xl font-semibold tracking-tight">
                                Crear una cuenta
                            </h1>
                            <p className="text-sm text-muted-foreground">
                                Ingrese su correo electronico para crear una
                                cuenta
                            </p>
                        </div>
                        <UserAuthForm />
                        <p className="px-8 text-center text-sm text-muted-foreground">
                            Al hacer clic en continuar, aceptas nuestros{" "}
                            <Link
                                href="/terms"
                                className="underline underline-offset-4 hover:text-primary"
                            >
                                Terminos de Servicio
                            </Link>{" "}
                            y{" "}
                            <Link
                                href="/privacy"
                                className="underline underline-offset-4 hover:text-primary"
                            >
                                Politica de Privacidad
                            </Link>
                            .
                        </p>
                    </div>
                </div>
            </div>
        </>
    )
}
