
namespace streeva.examples.profiler
{
    using System;
    using System.Threading;
    using Microsoft.Extensions.DependencyInjection;
    using StructureMap;
    using System.Reactive;

    class Program
    {
        static void Main(string[] args)
        {
            var services = new ServiceCollection();
            var container = new Container();
            container.Configure(config =>
            {
                config.For<IDataSource>().Use<Alphabet>();
                config.Populate(services);
            });

            var dataSource = container.GetInstance<IDataSource>();

            using (var token = dataSource.DataFeed.Subscribe(Observer.Create<Data>(OnNext, OnError, OnCompleted)))
            {
                Console.WriteLine("Press <Esc> to exit");
                while (!(Console.KeyAvailable && Console.ReadKey(true).Key == ConsoleKey.Escape))
                {
                    Thread.Sleep(TimeSpan.FromMilliseconds(100));
                }
            }

            ((Alphabet)dataSource).Dispose();
        }

        public static void OnCompleted()
        {
            Console.WriteLine("Done Signalled");
            Environment.Exit(0);
        }

        public static void OnError(Exception error) => Console.Error.WriteLine(error.Message);

        public static void OnNext(Data value) => Console.Write($"{value.Item} ");
    }
}
