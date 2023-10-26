export interface Tag {
    ID?: number
    name: string
}

export interface Job {
    ID?: number
    title: string
    abbr: string
    company: string
    description: string
    gradient?: string
    tags: Tag[]
    color: string
}
