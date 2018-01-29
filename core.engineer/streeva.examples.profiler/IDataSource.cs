
namespace streeva.examples.profiler
{
    using System;

    public interface IDataSource
    {
        IObservable<Data> DataFeed { get; }
    }
}