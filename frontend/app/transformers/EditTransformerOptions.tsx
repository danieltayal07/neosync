import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet';
import {
  CustomTransformer,
  Transformer,
  TransformerConfig,
} from '@/neosync-api-client/mgmt/v1alpha1/transformer_pb';
import { Cross2Icon, Pencil1Icon } from '@radix-ui/react-icons';
import { ReactElement, useEffect, useRef, useState } from 'react';
import CreditCardTransformerForm from './Sheetforms/CreditCardTransformerForm';
import EmailTransformerForm from './Sheetforms/EmailTransformerForm';
import FirstNameTransformerForm from './Sheetforms/FirstnameTransformerForm';
import FullNameTransformerForm from './Sheetforms/FullnameTransformerForm';
import GenderTransformerForm from './Sheetforms/GenderTransformerForm';
import IntPhoneNumberTransformerForm from './Sheetforms/IntPhoneNumberTransformerForm';
import LastNameTransformerForm from './Sheetforms/LastnameTransformerForm';
import PhoneNumberTransformerForm from './Sheetforms/PhoneNumberTransformerForm';
import RandomFloatTransformerForm from './Sheetforms/RandomFloatTransformerForm';
import RandomIntTransformerForm from './Sheetforms/RandomIntTransformerForm';
import RandomStringTransformerForm from './Sheetforms/RandomStringTransformerForm';
import UuidTransformerForm from './Sheetforms/UuidTransformerForm';

interface Props {
  transformer: CustomTransformer | undefined;
  index: number;
}

export default function EditTransformerOptions(props: Props): ReactElement {
  const { transformer, index } = props;

  const [isSheetOpen, setIsSheetOpen] = useState(false);
  const sheetRef = useRef<HTMLDivElement | null>(null);

  // handles click outside of sheet so that it closes correctly
  useEffect(() => {
    const handleOutsideClick = (event: MouseEvent) => {
      if (
        sheetRef.current &&
        !sheetRef.current.contains(event.target as Node)
      ) {
        setIsSheetOpen(false);
      }
    };

    if (isSheetOpen) {
      document.addEventListener('mousedown', handleOutsideClick);
    }

    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, [isSheetOpen]);

  // since component is in a controlled state, have to manually handle closing the sheet when the user presses escape
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        setIsSheetOpen!(false);
      }
    };
    window.addEventListener('keydown', handleKeyDown);
    // Clean up the event listener when the component unmounts
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, []);

  return (
    <Sheet open={isSheetOpen} onOpenChange={() => setIsSheetOpen(true)}>
      <SheetTrigger asChild>
        <Button
          variant="outline"
          size="sm"
          disabled={!transformer}
          onClick={() => setIsSheetOpen(true)}
          className="ml-auto hidden h-[36px] lg:flex"
        >
          <Pencil1Icon />
        </Button>
      </SheetTrigger>
      <SheetContent className="w-[800px]" ref={sheetRef}>
        <SheetHeader>
          <div className="flex flex-row justify-between w-full">
            <div className="flex flex-col space-y-2">
              <div className="flex flex-row gap-2">
                <SheetTitle>{transformer?.name}</SheetTitle>
                <Badge variant="outline">{transformer?.type}</Badge>
              </div>
              <SheetDescription>{transformer?.description}</SheetDescription>
            </div>
            <Button variant="ghost" onClick={() => setIsSheetOpen(false)}>
              <Cross2Icon className="h-4 w-4" />
            </Button>
          </div>
          <Separator />
        </SheetHeader>
        <div className="pt-8">
          {transformer &&
            handleTransformerForm(transformer, index, setIsSheetOpen)}
        </div>
      </SheetContent>
    </Sheet>
  );
}

function handleTransformerForm(
  transformer: CustomTransformer,
  index?: number,
  setIsSheetOpen?: (val: boolean) => void
): ReactElement {
  switch (transformer.source) {
    case 'email':
      return (
        <EmailTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'uuid':
      return (
        <UuidTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'first_name':
      return (
        <FirstNameTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'last_name':
      return (
        <LastNameTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'full_name':
      return (
        <FullNameTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'phone_number':
      return (
        <PhoneNumberTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'int_phone_number':
      return (
        <IntPhoneNumberTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'random_string':
      return (
        <RandomStringTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'random_int':
      return (
        <RandomIntTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'random_float':
      return (
        <RandomFloatTransformerForm
          index={index}
          transformer={transformer}
          setIsSheetOpen={setIsSheetOpen}
        />
      );
    case 'gender':
      return (
        <GenderTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    case 'credit_card':
      return (
        <CreditCardTransformerForm
          index={index}
          setIsSheetOpen={setIsSheetOpen}
          transformer={transformer}
        />
      );
    default:
      <div>No transformer component found</div>;
  }
  return <div></div>;
}

interface TransformerMetadata {
  name: string;
  description: string;
  type: string;
}

export function handleTransformerMetadata(
  value: string | undefined
): TransformerMetadata {
  const tEntries: Record<string, TransformerMetadata>[] = [
    {
      email: {
        name: 'Email',
        description: 'Anonymizes or generates a new email.',
        type: 'string',
      },
    },
    {
      phone_number: {
        name: 'Phone Number',
        description:
          'Anonymizes or generates a new phone number. The default format is <XXX-XXX-XXXX>.',
        type: 'string',
      },
    },
    {
      int_phone_number: {
        name: 'Int64 Phone Number',
        description:
          'Anonymizes or generates a new phone number of type int64 with a default length of 10.',
        type: 'int64',
      },
    },
    {
      first_name: {
        name: 'First Name',
        description: 'Anonymizes or generates a new first name.',
        type: 'string',
      },
    },
    {
      last_name: {
        name: 'Last Name',
        description: 'Anonymizes or generates a new last name.',
        type: 'string',
      },
    },
    {
      full_name: {
        name: 'Full Name',
        description:
          'Anonymizes or generates a new full name consisting of a first and last name.',
        type: 'string',
      },
    },
    {
      uuid: {
        name: 'UUID',
        description: 'Generates a new UUIDv4 id.',
        type: 'uuid',
      },
    },
    {
      passthrough: {
        name: 'Passthrough',
        description:
          'Passes the input value through to the desination with no changes.',
        type: 'passthrough',
      },
    },
    {
      null: {
        name: 'Null',
        description: 'Inserts a <null> string instead of the source value.',
        type: 'null',
      },
    },
    {
      random_string: {
        name: 'Random String',
        description:
          'Creates a randomly ordered alphanumeric string with a default length of 10 unless the String Length or Preserve Length parameters are defined.',
        type: 'string',
      },
    },
    {
      random_bool: {
        name: 'Random Bool',
        description: 'Generates a boolean value at random.',
        type: 'bool',
      },
    },
    {
      random_int: {
        name: 'Random Integer',
        description:
          'Generates a random integer value with a default length of 4 unless the Integer Length or Preserve Length paramters are defined. .',
        type: 'int64',
      },
    },
    {
      random_float: {
        name: 'Random Float',
        description:
          'Generates a random float value with a default length of <XX.XXX>.',
        type: 'float',
      },
    },
    {
      gender: {
        name: 'Gender',
        description:
          'Randomly generates one of the following genders: female, male, undefined, nonbinary.',
        type: 'string',
      },
    },
    {
      utc_timestamp: {
        name: 'UTC Timestamp',
        description: 'Randomly generates a UTC timestamp.',
        type: 'time',
      },
    },
    {
      unix_timestamp: {
        name: 'Unix Timestamp',
        description: 'Randomly generates a Unix timestamp.',
        type: 'int64',
      },
    },
    {
      street_address: {
        name: 'Street Address',
        description:
          'Randomly generates a street address in the format: {street_num} {street_addresss} {street_descriptor}. For example, 123 Main Street.',
        type: 'string',
      },
    },
    {
      city: {
        name: 'City',
        description:
          'Randomly selects a city from a list of predefined US cities.',
        type: 'string',
      },
    },
    {
      zipcode: {
        name: 'Zip Code',
        description:
          'Randomly selects a zip code from a list of predefined US cities.',
        type: 'string',
      },
    },
    {
      state: {
        name: 'State',
        description:
          'Randomly selects a US state and returns the two-character state code.',
        type: 'string',
      },
    },
    {
      full_address: {
        name: 'Full Address',
        description:
          'Randomly generates a street address in the format: {street_num} {street_addresss} {street_descriptor} {city}, {state} {zipcode}. For example, 123 Main Street Boston, Massachusetts 02169. ',
        type: 'string',
      },
    },
    {
      credit_card: {
        name: 'Credit Card',
        description:
          'Randomly generates a 16 digit credit card number with an option to generate a luhn valid credit card number.',
        type: 'int64',
      },
    },
    {
      sha256_hash: {
        name: 'SHA256 Hash',
        description:
          'SHA256 hashes the input value and returns back a string representation of the hash.',
        type: 'string',
      },
    },
  ];

  const def = {
    default: {
      name: 'Passthrough',
      description: 'Passthrough',
      type: 'passthrough',
    },
  };

  if (!value) {
    return def.default;
  }
  const res = tEntries.find((item) => item[value]);

  return res ? res[value] : def.default;
}

// merge system into custom and add in additional metadata fields for system transformers
// to fit into the custom transformers interface
export function MergeSystemAndCustomTransformers(
  system: Transformer[],
  custom: CustomTransformer[]
): CustomTransformer[] {
  let merged: CustomTransformer[] = [...custom];

  system.map((st) => {
    const cf = {
      config: {
        case: st.config?.config.case,
        value: st.config?.config.value,
      },
    };

    const newCt = new CustomTransformer({
      name: handleTransformerMetadata(st.value).name,
      description: handleTransformerMetadata(st.value).description,
      type: handleTransformerMetadata(st.value).type,
      source: st.value,
      config: cf as TransformerConfig,
    });

    merged.push(newCt);
  });

  return merged;
}
