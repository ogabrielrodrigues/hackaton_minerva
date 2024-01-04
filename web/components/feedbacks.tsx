import { Feedback as FeedbackType } from '@/types/feedback'
import { api_url } from '@/config/api'
import { GetEmployee } from '@/hooks/get_employee'
import { Feedback } from './feedback'
import { GetToken } from '@/hooks/get_token'
import { FilterForm } from './forms/filter_form'
import { GetFilter } from '@/hooks/get_filter'

export async function Feedbacks() {
  const { token } = await GetToken()
  const { employee } = await GetEmployee()
  const { filter } = await GetFilter()

  const url = filter
    ? `${api_url}/feedback/filter?unit=${employee?.unit}&sector=${filter}`
    : `${api_url}/feedback/filter?unit=${employee?.unit}`

  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const feedbacks = (await response.json()) as FeedbackType[]

  return (
    <>
      <FilterForm />
      <ul className="flex flex-col gap-2 overflow-y-scroll last:mb-12">
        {feedbacks?.length != 0 ? (
          feedbacks
            ?.filter(
              feedback => feedback.employee_registry != employee?.registry
            )
            ?.map(feedback => (
              <Feedback
                key={feedback.id}
                employee_registry={feedback.employee_registry}
                content={feedback.content}
                sent_at={feedback.sent_at}
              />
            ))
        ) : (
          <p className="self-center mt-16 font-medium text-primary">
            Nenhum feedback disponível!
          </p>
        )}
      </ul>
    </>
  )
}
