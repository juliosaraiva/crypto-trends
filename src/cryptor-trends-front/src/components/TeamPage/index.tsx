"use client"
import { Card, CardBody, Divider, Image, Link, Spacer } from '@nextui-org/react';
import { GithubLogo, LinkedinLogo } from '@phosphor-icons/react';
import React from 'react';

const TeamPage: React.FC = () => {
  const teamMembers = [
    {
      id: 1,
      name: 'Edson Costa',
      title: 'Software Engineer',
      imageUrl: 'https://github.com/ecsistem.png',
      githubUrl: 'https://github.com/ecsistem',
      linkedinUrl: 'https://www.linkedin.com/in/edsoncostadev',
    },
    {
      id: 2,
      name: 'Fillipe Nascimento',
      title: 'Software Engineer',
      imageUrl: 'https://github.com/linkinn.png',
      githubUrl: 'https://github.com/linkinn',
      linkedinUrl: 'https://www.linkedin.com/in/fillipi-nascimento-35128070/',
    },
    {
      id: 3,
      name: 'Julio Saraiva',
      title: 'Software Engineer',
      imageUrl: 'https://github.com/juliosaraiva.png',
      githubUrl: 'https://github.com/juliosaraiva',
      linkedinUrl: 'https://www.linkedin.com/in/ojuliosaraiva',
    },
  ];

  return (
    <div className="container mx-auto">
      <h1 className="text-3xl font-bold mb-8 text-center">Meet the Team</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 justify-center">
        {teamMembers.map((member) => (
          <div key={member.id}>
            <Card className="w-full">
              <CardBody>
                <div className="flex flex-col items-center gap-4">
                  <Image
                    alt="Card background"
                    className="object-cover rounded-full"
                    src={member.imageUrl}
                    width={200}
                  />
                  <div className="text-center">
                    <h5 className="font-bold text-lg">{member.name}</h5>
                    <Spacer y={1} />
                    <h5 className="text-sm">{member.title}</h5>
                  </div>
                </div>
                <Divider className="my-4" />
                <div className="flex justify-center gap-4">
                  <Link href={member.linkedinUrl} target="_blank" className="text-primary flex items-center">
                    <LinkedinLogo size={24} />
                  </Link>
                  <Link href={member.githubUrl} target="_blank" className="text-primary flex items-center">
                    <GithubLogo size={24} />
                  </Link>
                </div>
              </CardBody>
            </Card>
          </div>
        ))}
      </div>
    </div>
  );
};

export default TeamPage;
