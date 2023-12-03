'use client';

import FormError from '@/components/FormError';
import OverviewContainer from '@/components/containers/OverviewContainer';
import PageHeader from '@/components/headers/PageHeader';
import { useAccount } from '@/components/providers/account-provider';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from '@/components/ui/command';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectTrigger } from '@/components/ui/select';
import { toast } from '@/components/ui/use-toast';
import { useGetSystemTransformers } from '@/libs/hooks/useGetSystemTransformers';
import { cn } from '@/libs/utils';
import {
  CreateUserDefinedTransformerRequest,
  CreateUserDefinedTransformerResponse,
  SystemTransformer,
  TransformerConfig,
} from '@/neosync-api-client/mgmt/v1alpha1/transformer_pb';
import { getErrorMessage } from '@/util/util';
import { yupResolver } from '@hookform/resolvers/yup';
import { CheckIcon } from '@radix-ui/react-icons';
import { useRouter } from 'next/navigation';
import { ReactElement, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { handleCustomTransformerForm } from './UserDefinedTransformerForms/HandleCustomTransformersForm';
import {
  CREATE_USER_DEFINED_TRANSFORMER_SCHEMA,
  CreateUserDefinedTransformerSchema,
} from './schema';

export default function NewTransformer(): ReactElement {
  const [base, setBase] = useState<SystemTransformer>(
    new SystemTransformer({})
  );
  const [openBaseSelect, setOpenBaseSelect] = useState(false);

  const form = useForm<CreateUserDefinedTransformerSchema>({
    resolver: yupResolver(CREATE_USER_DEFINED_TRANSFORMER_SCHEMA),
    defaultValues: {
      name: '',
      source: '',
      type: '',
      config: {},
      description: '',
    },
  });

  const router = useRouter();
  const { account } = useAccount();

  async function onSubmit(
    values: CreateUserDefinedTransformerSchema
  ): Promise<void> {
    if (!account) {
      return;
    }
    try {
      const transformer = await createNewTransformer(account.id, values);
      toast({
        title: 'Successfully created transformer!',
        variant: 'success',
      });
      if (transformer.transformer?.id) {
        router.push(`/transformers/${transformer.transformer?.id}`);
      } else {
        router.push(`/transformers`);
      }
    } catch (err) {
      console.error(err);
      toast({
        title: 'Unable to create transformer',
        description: getErrorMessage(err),
        variant: 'destructive',
      });
    }
  }

  const { data } = useGetSystemTransformers();
  const transformers = data?.transformers ?? [];

  return (
    <OverviewContainer
      Header={<PageHeader header="Create a new Transformer" />}
      containerClassName="px-12 md:px-24 lg:px-32"
    >
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="source"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Source Transformer</FormLabel>
                <FormDescription>
                  The system transformer to clone.
                </FormDescription>
                <FormControl>
                  <Select
                    open={openBaseSelect}
                    onOpenChange={setOpenBaseSelect}
                  >
                    <SelectTrigger>
                      {base.name ? base.name : 'Select a transformer'}
                    </SelectTrigger>
                    <SelectContent>
                      <Command className="overflow-auto">
                        <CommandInput placeholder="Search transformers..." />
                        <CommandEmpty>No transformers found.</CommandEmpty>
                        <CommandGroup className="overflow-auto h-[200px]">
                          {transformers.map((t, index) => (
                            <CommandItem
                              key={`${t.source}-${index}`}
                              onSelect={(value: string) => {
                                const selectedTransformer = transformers.find(
                                  (item) => item.source == value
                                );

                                field.onChange(selectedTransformer?.source);
                                const selectedValues = {
                                  config: {
                                    case: selectedTransformer?.config?.config
                                      .case,
                                    value:
                                      selectedTransformer?.config?.config
                                        .value || {},
                                  },
                                };
                                form.setValue('config', selectedValues);
                                form.setValue(
                                  'type',
                                  selectedTransformer?.dataType ?? ''
                                );
                                setBase(
                                  selectedTransformer ??
                                    new SystemTransformer({})
                                );
                                setOpenBaseSelect(false);
                              }}
                              value={t.source}
                            >
                              <CheckIcon
                                className={cn(
                                  'mr-2 h-4 w-4',
                                  base.name == t.name
                                    ? 'opacity-100'
                                    : 'opacity-0'
                                )}
                              />
                              {t.name}
                            </CommandItem>
                          ))}
                        </CommandGroup>
                      </Command>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          {form.getValues('source') && (
            <div>
              <Controller
                control={form.control}
                name="name"
                render={({ field: { onChange, ...field } }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormDescription>
                      The unique name of the transformer.
                    </FormDescription>
                    <FormControl>
                      <Input
                        placeholder="Transformer Name"
                        {...field}
                        onChange={async ({ target: { value } }) => {
                          onChange(value);
                          await form.trigger('name');
                        }}
                      />
                    </FormControl>
                    <FormError
                      errorMessage={form.formState.errors.name?.message ?? ''}
                    />
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="pt-10">
                <FormField
                  control={form.control}
                  name="description"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Description</FormLabel>
                      <FormDescription>
                        The Transformer decription.
                      </FormDescription>
                      <FormControl>
                        <Input
                          placeholder="Transformer description"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </div>
          )}
          <div>{handleCustomTransformerForm(form.getValues('source'))}</div>
          <div className="flex flex-row justify-end">
            <Button type="submit" disabled={!form.formState.isValid}>
              Next
            </Button>
          </div>
        </form>
      </Form>
    </OverviewContainer>
  );
}

async function createNewTransformer(
  accountId: string,
  formData: CreateUserDefinedTransformerSchema
): Promise<CreateUserDefinedTransformerResponse> {
  const body = new CreateUserDefinedTransformerRequest({
    accountId: accountId,
    name: formData.name,
    description: formData.description,
    type: formData.type,
    source: formData.source,
    transformerConfig: formData.config as TransformerConfig,
  });

  const res = await fetch(`/api/transformers/user-defined`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: JSON.stringify(body),
  });
  if (!res.ok) {
    const body = await res.json();
    throw new Error(body.message);
  }
  return CreateUserDefinedTransformerResponse.fromJson(await res.json());
}
