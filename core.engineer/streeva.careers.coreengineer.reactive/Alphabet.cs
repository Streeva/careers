
namespace streeva.careers.coreengineer.reactive
{
    using System;
    using System.Reactive.Linq;
    using System.Reactive.Subjects;
    using System.Threading;

    public class Alphabet : IDataSource, IDisposable
    {
        #region public

        public Alphabet()
        {
            this.timer = new Timer(this.PublishLetter, null, TimeSpan.FromSeconds(1), TimeSpan.FromSeconds(1));
        }

        public IObservable<Data> DataFeed => this.dataSubject.AsObservable<Data>();

        public void Dispose()
        {
            this.dataSubject.OnCompleted();
            this.timer.Dispose();
        }

        #endregion

        #region private

        private void PublishLetter(object state)
        {

            this.dataSubject.OnNext(new Data(this.data[this.index++].ToString()));
            if (this.index >= this.data.Length)
            {
                this.index = 0;
                return;
            }
        }

        private Subject<Data> dataSubject = new Subject<Data>();

        private string data = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        private int index;

        private Timer timer;

        #endregion
    }
}