'use server'

import { revalidatePath } from 'next/cache'
import { cookies } from 'next/headers'

export async function filterAction(_: any, fd: FormData) {
  try {
    const sector_filter = fd.get('sector')

    if (sector_filter?.toString().trim() == '') {
      cookies().delete('feedback_filter')
      revalidatePath('/')
      return { filtered: false }
    }

    if (!sector_filter) {
      return { filtered: false }
    }

    cookies().set({
      name: 'feedback_filter',
      value: sector_filter.toString(),
      path: '/',
      maxAge: 86400,
    })

    revalidatePath('/')
    return { filtered: true }
  } catch (err) {
    return { filtered: false }
  }
}
