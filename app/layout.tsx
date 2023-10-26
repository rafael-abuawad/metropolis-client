import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"
import { ThemeProvider } from "@/components/theme-provider"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
    title: "Metropolis",
    description:
        "La principal plataforma de empleo en el país. Nuestro portal de ofertas laborales y oportunidades de empleo cuenta con una amplia y completa cobertura en todo el territorio boliviano.",
}

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
            <ThemeProvider
                attribute="class"
                defaultTheme="system"
                enableSystem
                disableTransitionOnChange
            >
                <body className={inter.className}>{children}</body>
            </ThemeProvider>
        </html>
    )
}
