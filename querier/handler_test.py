import os
import unittest
from unittest.mock import patch


class TestQuerier(unittest.TestCase):
    def __init__(self, methodName):
        super().__init__(methodName)
        os.environ['REGION'] = 'unit-test'
        os.environ['ES_HOST'] = 'http://unit-test'

    @patch('logging.Logger.info')
    @patch('elasticsearch.Elasticsearch.search')
    def test_querier(self, search, logger):
        from handler import querier

        def mock_elastic_search(index=None, body=None):
            self.assertEqual(index, 'events')
            self.assertEqual(body['size'], 0)
            return {"aggregations": {
                "num_users": {"value": 111},
                "num_events": {"value": 222}
            }}

        def mock_logger(msg=None, extra=None):
            self.assertEqual(msg, 'querier metric')
            self.assertEqual(extra['num_users'], 111)
            self.assertEqual(extra['num_events'], 222)

        search.side_effect = mock_elastic_search
        logger.side_effect = mock_logger

        querier(None, None)

        self.assertEqual(search.called, True)
        self.assertEqual(logger.called, True)


if __name__ == '__main__':
    unittest.main()
